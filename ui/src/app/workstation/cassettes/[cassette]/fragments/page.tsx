'use client'

import { Fragment, useEffect, useState } from "react";
import { useRouter } from "next/navigation";
import Image from "next/image";
import { IoIosAdd } from "react-icons/io";
import { FiChevronLeft, FiRotateCcw } from "react-icons/fi";
import { PiPushPinThin } from "react-icons/pi";
import { MdArrowBackIosNew } from "react-icons/md";

import useFragment from "@/hooks/useFragment";
import { FragmentCard, Input, LoadingSpinner, MediumButton, PageActionMenu, SlideOver } from "@/components";
import { AiOutlinePlus } from "react-icons/ai";
import { updateAppState } from "@/store/slices/appStateSlice";
import { menuPath } from '@/utils/app.utils'
import { positions } from "@/utils/dom.utils";

export default function CassetteFragments({ params }: { params: { cassette: string } }) {

    const { name, setName, fragments, onSubmit, onDelete, onGetAll, laoding, cassettesSelectList, setId, onUpdate } = useFragment()
    const [createSlideOverState, setCreateSlideOverState] = useState(false)
    const [updateSlideOverState, setUpdateSlideOverState] = useState(false)
    const router = useRouter()

    const onCreateSlideOver = () => {
        setCreateSlideOverState(!createSlideOverState);
    }

    const onUpdateSlideOver = (cassetteId: string, cassetteName: string) => {
        setName(cassetteName)
        setId(cassetteId)
        setUpdateSlideOverState(!updateSlideOverState);
    }

    const onSaveUpdate = async () => {
        await onUpdate().then((result) => result ? onUpdateSlideOver('', '') : null)
    }


    const onRefreshList = () => {
        onGetAll(params.cassette);
    }

    const onFetchBranches = async (fragmentId: string) => {
        router.push(`/workstation/cassettes/fragments/${fragmentId}/branches`)
    }

    const onBackToCassettes = async () => {
        router.back()
    }

    useEffect(() => {
        onGetAll(params.cassette);
    }, [])

    return (
        <Fragment>
            <SlideOver openState={createSlideOverState} openHandler={onCreateSlideOver} title="Create New Fragment">
                <form>
                    <div className="flex flex-col">
                        <div>
                            <Input
                                id="name"
                                label="Fragment Name"
                                type="text"
                                extra="mb-3"
                                onChange={setName}
                                value={name}
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
            <SlideOver openState={updateSlideOverState} openHandler={onUpdateSlideOver} title="Edit Fragment">
                <form>
                    <div className="flex flex-col">
                        <div>
                            <Input
                                id="name_update"
                                label="Fragment Name"
                                type="text"
                                extra="mb-3"
                                onChange={setName}
                                value={name}
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
                    text: 'Back to cassettes',
                    position: positions.TopLeft,
                    handler: onBackToCassettes,
                    classes: 'bg-gray-100 px-4 py-2 text-base font-medium text-slate-500 transition duration-200 hover:bg-gray-200 active:bg-gray-300 dark:bg-white/10 dark:text-white dark:hover:bg-white/20 dark:active:bg-white/30 '
                }
            ]} filter={[]} />
            <div>
                {
                    laoding ? <LoadingSpinner width={28} height={28} /> : (
                        fragments.length > 0 ? (fragments.map((fragment, i) => {
                            return (
                                <FragmentCard key={i} fragment={fragment} onDelete={onDelete} onFetchBranches={onFetchBranches} onUpdateOpen={onUpdateSlideOver}/>
                            )
                        })) :
                            <Fragment>
                                <p className="text-slate-500 dark:text-amber-50 text-center">{`This cassette is empty.Create a fragment and record on it.`}</p>
                            </Fragment>
                    )
                }
            </div>
        </Fragment >
    );

}