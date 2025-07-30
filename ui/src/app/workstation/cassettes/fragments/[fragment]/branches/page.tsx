'use client'

import { Fragment, useEffect, useState } from "react";
import { useRouter } from "next/navigation";
import { FiRotateCcw } from "react-icons/fi";
import { IoIosAdd } from "react-icons/io";
import { AiOutlinePlus } from "react-icons/ai";

import useBranch from "@/hooks/useBranch";
import { BranchCard, Input, LoadingSpinner, MediumButton, PageActionMenu, SlideOver } from "@/components";
import { menuPath } from '@/utils/app.utils'
import { positions } from "@/utils/dom.utils";
import { MdArrowBackIosNew } from "react-icons/md";


export default function Branches({ params }: { params: { fragment: string } }) {

    const { name, setName, email, setEmail, description, setDescription, branches, onSubmit, onDeleteRequest, onGetAll, laoding, setId, onUpdate } = useBranch()
    const [createSlideOverState, setCreateSlideOverState] = useState(false)
    const [updateSlideOverState, setUpdateSlideOverState] = useState(false)
    const router = useRouter()

    const onCreateSlideOver = () => {
        setCreateSlideOverState(!createSlideOverState);
    }

    const onUpdateSlideOver = (cassetteId: string, cassetteName: string) => {
        setId(cassetteId)
        setName(cassetteName)
        setUpdateSlideOverState(!updateSlideOverState);
    }

    const onSaveUpdate = async () => {
        await onUpdate().then((result) => result ? onUpdateSlideOver('', '') : null)
    }

    const onRefreshList = () => {
        onGetAll(params.fragment);
    }

    const onFetchTimes = async (branchId: string) => {
        router.push(`/workstation/cassettes/fragments/branches/${branchId}/times`)
    }

    const onBackToFragments = async () => {
        router.back()
    }

    useEffect(() => {
        onGetAll(params.fragment);
    }, [])

    return (
        <Fragment>
            <SlideOver openState={createSlideOverState} openHandler={onCreateSlideOver} title="Create New Branch">
                <form>
                    <div className="flex flex-col">
                        <div>
                            <Input
                                id="name"
                                label="Branch Name"
                                type="text"
                                extra="mb-3"
                                onChange={setName}
                                value={name}
                            />
                        </div>
                        <div>
                            <Input
                                id="email"
                                label="User Email"
                                type="email"
                                extra="mb-3"
                                onChange={setEmail}
                                value={email}
                            />
                        </div>
                        <div>
                            <MediumButton type="button" text="Create" extraClasses="w-full" onClick={() => onSubmit()} roundedSize="lg">
                                <AiOutlinePlus />
                            </MediumButton>
                        </div>
                    </div>
                </form>
            </SlideOver>
            <SlideOver openState={updateSlideOverState} openHandler={onUpdateSlideOver} title="Edit Branch">
                <form>
                    <div className="flex flex-col">
                        <div>
                            <Input
                                id="name_update"
                                label="Branch Name"
                                type="text"
                                extra="mb-3"
                                onChange={setName}
                                value={name}
                            />
                        </div>
                        <div>
                            <Input
                                id="email_update"
                                label="User Email"
                                type="email"
                                extra="mb-3"
                                onChange={setEmail}
                                value={email}
                            />
                        </div>
                        <div>
                            <MediumButton type="button" text="Save" extraClasses="w-full" onClick={() => onSaveUpdate()} roundedSize="lg">
                                <AiOutlinePlus />
                            </MediumButton>
                        </div>
                    </div>
                </form>
            </SlideOver>
            <PageActionMenu actions={[
                {
                    icon: <IoIosAdd />,
                    text: 'Create',
                    position: positions.TopRight,
                    handler: onCreateSlideOver,
                    classes: 'rounded-lg dark:bg-white-800 dark:bg-emerald-800 bg-blue-600 px-3 text-white hover:bg-brand-600 active:bg-brand-700  dark:hover:bg-brand-300 dark:active:bg-brand-200  rounded-lg w-full p-2 text-xs text-base font-medium text-white transition duration-200 hover:bg-brand-600 active:bg-brand-700 dark:bg-brand-400 dark:text-white dark:hover:bg-brand-300 dark:active:bg-brand-200 '
                },
                {
                    icon: <FiRotateCcw />,
                    text: 'Refresh',
                    position: positions.TopRight,
                    handler: onRefreshList,
                    classes: 'rounded-lg dark:bg-white-800 bg-teal-500 px-3 text-white hover:bg-brand-600 active:bg-brand-700  dark:hover:bg-brand-300 dark:active:bg-brand-200  rounded-lg w-full p-2 text-xs text-base font-medium text-white transition duration-200 hover:bg-brand-600 active:bg-brand-700 dark:bg-brand-400 dark:text-white dark:hover:bg-brand-300 dark:active:bg-brand-200 '
                },
                {
                    icon: <MdArrowBackIosNew />,
                    text: 'Back to fragments',
                    position: positions.TopLeft,
                    handler: onBackToFragments,
                    classes: 'bg-gray-100 px-4 py-2 text-base font-medium text-slate-500 transition duration-200 hover:bg-gray-200 active:bg-gray-300 dark:bg-white/10 dark:text-white dark:hover:bg-white/20 dark:active:bg-white/30 '
                }
            ]} filter={[]} />

            {laoding ? <LoadingSpinner width={28} height={28} /> : (
                branches.length > 0 ? (branches.map((branch, i) => {
                    return (
                        <BranchCard key={i} branch={branch} onDeleteRequest={onDeleteRequest} onFetchTimes={onFetchTimes} onUpdateOpen={onUpdateSlideOver} />
                    )
                })) :
                    <Fragment>
                        <p className="text-slate-500 dark:text-amber-50 text-center">{`You haven't created any branch for this fragment yet.`}</p>
                    </Fragment>
            )}
        </Fragment>
    );

}