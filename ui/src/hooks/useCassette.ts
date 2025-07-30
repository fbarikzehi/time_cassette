import { CassetteModel } from '@/models/Cassette'
import { useState } from 'react'
import toast from 'react-hot-toast';
import { DispatchRequest, RequestMethods } from "../utils/api.utils";
import { ApiUrls } from "@/utils/api.url.utils";

export default function useCassette() {
    const cassetteList: CassetteModel[] = [];
    const [cassettes, setCassettes] = useState(cassetteList)
    const [id, setId] = useState('')
    const [name, setName] = useState('')
    const [laoding, setLoading] = useState(false)

    const onGetAll = async () => {
        setLoading(true)
        await DispatchRequest(RequestMethods.Get, ApiUrls.Cassette.Cassettes, null).then(response => {
            response.meta.result && response.data ? setCassettes(response.data.reverse()) : setCassettes([]);
            setLoading(false)
            return response.meta.result;
        })
    }

    const onSubmit = async () => {
        if (!name) {
            toast.error("Name is required")
            return false;
        }
        const toastId = toast.loading('Creating a new cassette...');
        await DispatchRequest(RequestMethods.Post, ApiUrls.Cassette.Create, { name }).then(response => {
            if (response.meta.result) {
                setName('')
                onGetAll()
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
        const toastId = toast.loading('Editing cassette...');
        const response = await DispatchRequest(RequestMethods.Put, ApiUrls.Cassette.Update, { id, name }).then(response => {
            if (response.meta.result) {
                let index = cassettes.findIndex(c => c.Id == id);
                let oldOne = cassettes[index]
                cassettes.splice(index, 1, {
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
        await DispatchRequest(RequestMethods.Delete, ApiUrls.Cassette.Delete, { id }).then((response) => {
            if (response.meta.result) {
                onGetAll()
                toast.success(response.meta.messages)
            }
            else {
                toast.error(response.meta.messages)
            }
        })

    }

    return { name, setName, setId, cassettes, onSubmit, onDelete, onUpdate, onGetAll, laoding }
}
