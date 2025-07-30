import { BranchModel } from '@/models/Branch'
import { useState } from 'react'
import toast from 'react-hot-toast';

import { DispatchRequest, RequestMethods } from "../utils/api.utils";
import { ApiUrls } from "../utils/api.url.utils";

export default function useBranch() {

    const branchList: BranchModel[] = [];
    const [branches, setBranches] = useState(branchList)
    const [id, setId] = useState('')
    const [name, setName] = useState('')
    const [email, setEmail] = useState('')
    const [description, setDescription] = useState('')
    const [secretCode, setSecretCode] = useState('')
    const [fragmentId, setFragmentId] = useState('')

    const [laoding, setLoading] = useState(false)

    const onGetAll = async (fragmentId: string) => {
        if (fragmentId) setFragmentId(fragmentId)
        setLoading(true)
        await DispatchRequest(RequestMethods.Get, ApiUrls.Fragment.Branches, { fragmentId }).then(response => {
            response.meta.result && response.data ? setBranches(response.data.reverse()) : setBranches([]);
            setLoading(false)
            return response.meta.result;
        })
    }

    const onSubmit = async () => {
        if (!fragmentId) {
            toast.error("Fragment is not loaded")
            return false;
        }
        if (!name) {
            toast.error("Name is required")
            return false;
        }
        if (!email) {
            toast.error("User Email is required")
            return false;
        }
        const toastId = toast.loading('Creating a new branch on this fragment...');
        await DispatchRequest(RequestMethods.Post, ApiUrls.Branch.Create, { name, fragmentId, handlerUserEmail: email }).then(response => {
            if (response.meta.result) {
                setName('')
                setEmail('')
                onGetAll(fragmentId)
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
        const toastId = toast.loading('Editing branch...');
        const response = await DispatchRequest(RequestMethods.Put, ApiUrls.Branch.Update, { id, name, fragmentId }).then(response => {
            if (response.meta.result) {
                let index = branches.findIndex(b => b.Id == id);
                let oldOne = branches[index]
                branches.splice(index, 1, {
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

    const onDeleteRequest = async (id: string) => {
        await DispatchRequest(RequestMethods.Post, ApiUrls.Branch.DeleteRequest, { id, description }).then((response) => {
            if (response.meta.result) {
                onGetAll(fragmentId)
                toast.success(response.meta.messages)
            }
            else {
                toast.error(response.meta.messages)
            }
        })
    }

    const onDeleteConfirm = async (id: string) => {
        if (!secretCode) {
            toast.error("Secret Code is required")
            return false;
        }
        await DispatchRequest(RequestMethods.Post, ApiUrls.Branch.DeleteConfirm, { id, secretCode }).then((response) => {
            if (response.meta.result) {
                onGetAll(fragmentId)
                toast.success(response.meta.messages)
            }
            else {
                toast.error(response.meta.messages)
            }
        })
    }

    // const onEmailSearch = async () => {
    //     setEmailLoading(true)
    //     await DispatchRequest(RequestMethods.Get, ApiUrls.User.SearchEmail, { email:searchEmail }).then(response => {
    //         response.meta.result && response.data ? setBranches(response.data.reverse()) : setBranches([]);
    //         setEmailLoading(false)
    //         return response.meta.result;
    //     })
    // }

    return { name, setName, setId, email, setEmail, description, setDescription, branches, onSubmit, onDeleteRequest, onDeleteConfirm, onUpdate, onGetAll, laoding, setSecretCode, fragmentId }
}
