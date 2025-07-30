import { TimeModel } from '@/models/Time'
import { useState } from 'react'
import toast from 'react-hot-toast';
import moment from "moment";
import { DispatchRequest, RequestMethods } from "../utils/api.utils";
import { ApiUrls } from "../utils/api.url.utils";

export default function useTime() {
    // import { useSelector } from 'react-redux';
    // const appState = useSelector((state: storeModel) => state.appState);

    const TimeList: TimeModel[] = [];
    const [times, setTimes] = useState(TimeList)

    const [id, setId] = useState('')
    const [duration, setDuration] = useState(0)
    // const [startDateTime, setStartDateTime] = useState(moment())
    const [branchId, setBranchId] = useState('')

    const [laoding, setLoading] = useState(false)

    const onGetAll = async (branchId: string) => {
        if (branchId) setBranchId(branchId)
        setLoading(true)
        await DispatchRequest(RequestMethods.Get, ApiUrls.Branch.Times, { branchId }).then(response => {
            if (response.meta.result && response.data) {
                let times = response.data.reverse()
                setTimes(times)

            } else { setTimes([]) };


            setLoading(false)
            return response.meta.result;
        })
    }

    const onInstantSubmit = async () => {
        if (!branchId) {
            toast.error("Branch is not loaded")
            return false;
        }
        const toastId = toast.loading('Creating a new time on this branch...');
        let nowDateTime = moment()
        let payload = { duration: duration, startDateTime: nowDateTime.format('YYYY-MM-DDTHH:mm:ss') + 'Z', endDateTime: nowDateTime.format('YYYY-MM-DDTHH:mm:ss') + 'Z', branchId: branchId }
        console.log(payload)
        await DispatchRequest(RequestMethods.Post, ApiUrls.Time.Create, payload).then(response => {
            if (response.meta.result) {
                setDuration(0)
                // setStartDateTime(moment())
                onGetAll(branchId)

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

    const onStop = async (id: string) => {
        if (!id) {
            toast.error("Time is not loaded")
            return false;
        }
        let nowDateTime = moment()
        await DispatchRequest(RequestMethods.Put, ApiUrls.Time.UpdateEnd, { id: id, endDateTime: nowDateTime.format('YYYY-MM-DDTHH:mm:ss') + 'Z' }).then(response => {
            if (response.meta.result) {
                onGetAll(branchId)
            }
            else {
                toast.error(response.meta.messages)
            }
            return response.meta.result;
        })
    }

    const onDelete = async (id: string) => {
        await DispatchRequest(RequestMethods.Delete, ApiUrls.Time.Delete, { id }).then((response) => {
            if (response.meta.result) {
                onGetAll(branchId)
            }
            else {
                toast.error(response.meta.messages)
            }
        })
    }

    const onDeleteAll = async () => {
        await DispatchRequest(RequestMethods.Delete, ApiUrls.Time.Delete, null).then((response) => {
            if (response.meta.result) {
                onGetAll(branchId)
                toast.success(response.meta.messages)
            }
            else {
                toast.error(response.meta.messages)
            }
        })
    }

    return { setDuration, duration, times, setTimes, onInstantSubmit, onDelete, onDeleteAll, onGetAll, onStop, laoding }
}
