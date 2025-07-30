'use client'

import { useState, useEffect, Fragment } from "react";
import { AiOutlinePlus } from "react-icons/ai";
import { useRouter } from 'next/navigation'

import { updateAppState } from "@/store/slices/appStateSlice";

import { CassetteCard, PageActionMenu, SlideOver, Input, MediumButton, ExtraSmallButton, LoadingSpinner } from "@/components";
import useCassette from '@/hooks/useCassette'
import { IoIosAdd } from "react-icons/io";
import { FiRotateCcw } from "react-icons/fi";
import { menuPath } from '@/utils/app.utils'
import { positions } from "@/utils/dom.utils";


export default function Cassettes() {
    const { name, setName, setId, cassettes, onSubmit, onDelete, onUpdate, onGetAll, laoding } = useCassette()
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

    const onFetchFragments = async (cassetteId: string) => {
        router.push(`/workstation/cassettes/${cassetteId}/fragments`)
    }

    const onRefreshList = () => {
        onGetAll();
        console.log(laoding)
    }

    useEffect(() => {
        onGetAll();
        console.log("Test actions")
    }, [])

    return (
        <Fragment>
            <SlideOver openState={createSlideOverState} openHandler={onCreateSlideOver} title="Create New Cassette">
                <form>
                    <div className="flex flex-col">
                        <div>
                            <Input
                                id="name_create"
                                label="Cassette Name"
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
            <SlideOver openState={updateSlideOverState} openHandler={onUpdateSlideOver} title="Edit Cassette">
                <form>
                    <div className="flex flex-col">
                        <div>
                            <Input
                                id="name_update"
                                label="Cassette Name"
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
            ]} filter={[]} />
            {laoding ? <LoadingSpinner width={28} height={28} /> : (
                cassettes.length > 0 ? (cassettes.map((cassette, i) => {
                    return (
                        <CassetteCard key={i} cassette={cassette} onFetchFragments={onFetchFragments} onDelete={onDelete} onUpdateOpen={onUpdateSlideOver} />
                    )
                })) :
                    <Fragment>
                        <p className="text-slate-500 dark:text-amber-50 text-center">{`Your cassette shelf is empty.`}</p>
                    </Fragment>
            )}
        </Fragment>
    );

}