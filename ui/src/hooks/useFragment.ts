import { FragmentModel } from '@/models/Fragment'
import { useState } from 'react'
import toast from 'react-hot-toast';

import { DispatchRequest, RequestMethods } from "../utils/api.utils";
import { ApiUrls } from "../utils/api.url.utils";
import { SeletcListOption } from '@/models/SelectListModel'

export default function useFragment() {

    const fragmentList: FragmentModel[] = [];
    const [fragments, setFragments] = useState(fragmentList)
    const [id, setId] = useState('')
    const [name, setName] = useState('')
    const [cassetteId, setCassetteId] = useState('')
    const [laoding, setLoading] = useState(false)
    const [cassettesSelectList, setCassettesSelectList] = useState<Array<SeletcListOption>>([])

    const onGetAll = async (cassetteId?: string) => {
        setLoading(true)
        if (cassetteId) {
            setCassetteId(cassetteId)
            await DispatchRequest(RequestMethods.Get, ApiUrls.Cassette.Fragments, { cassetteId }).then(response => {
                response.meta.result && response.data ? setFragments(response.data.reverse()) : setFragments([]);
                setLoading(false)
                return response.meta.result;
            })
        }
        setLoading(false)
        await DispatchRequest(RequestMethods.Get, ApiUrls.Cassette.Cassettes, null).then(response => {
            setCassettesSelectList([...[{ name: 'Cassettes', value: '', selected: true }], ...response.data.map((c: any) => {
                return { name: c.name, value: c.Id, selected: false }
            })])
            // console.log(cassettesSelectList)

        })
    }

    const onSubmit = async () => {
        if (!cassetteId) {
            toast.error("Cassette is not loaded")
            return false;
        }
        if (!name) {
            toast.error("Name is required")
            return false;
        }
        const toastId = toast.loading('Creating a new fragment on this cassette...');
        await DispatchRequest(RequestMethods.Post, ApiUrls.Fragment.Create, { name, cassetteId }).then(response => {
            if (response.meta.result) {
                setName('')
                onGetAll(cassetteId)
                toast.success(response.meta.messages, {
                    id: toastId,
                })
            }
            else {
                toast.error(response.meta.messages, {
                    id: toastId,
                })
            }
            return response.meta.result;
        })
    }


    const onUpdate = async () => {
        if (!name) {
            toast.error("Name is required")
            return false;
        }
        const toastId = toast.loading('Editing fragment...');
        const response = await DispatchRequest(RequestMethods.Put, ApiUrls.Fragment.Update, { id, name, cassetteId }).then(response => {
            if (response.meta.result) {
                let index = fragments.findIndex(f => f.Id == id);
                let oldOne = fragments[index]
                fragments.splice(index, 1, {
                    ...oldOne, ...{ name: name }
                })
                setId('')
                setName('')
                toast.success(response.meta.messages, {
                    id: toastId,
                })
            }
            else {
                toast.error(response.meta.messages, {
                    id: toastId,
                })
            }
            return response;
        })
        return response.meta.result;
    }

    const onDelete = async (id: string) => {
        await DispatchRequest(RequestMethods.Delete, ApiUrls.Fragment.Delete, { id }).then((response) => {
            if (response.meta.result) {
                onGetAll(cassetteId)
                toast.success(response.meta.messages)
            }
            else {
                toast.error(response.meta.messages)
            }
        })
    }

    return { name, setName, setId, fragments, onSubmit, onDelete, onUpdate, onGetAll, laoding, cassetteId, cassettesSelectList }
}
