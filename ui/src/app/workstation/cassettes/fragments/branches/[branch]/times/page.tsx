'use client'

import { Fragment, useEffect, useState } from "react";
import { FiRotateCcw } from "react-icons/fi";
import { TfiLayersAlt } from "react-icons/tfi";
import { AiOutlinePlus } from "react-icons/ai";
import { HiStopCircle } from "react-icons/hi2";

import useTime from "@/hooks/useTime";
import { TimeCard, Input, LoadingSpinner, MediumButton, PageActionMenu, SlideOver } from "@/components";
import { positions } from "@/utils/dom.utils";
import { MdArrowBackIosNew } from "react-icons/md";
import { RxPlay } from "react-icons/rx";
import { useRouter } from "next/navigation";
import moment, { duration } from "moment";


export default function Times({ params }: { params: { branch: string } }) {

    const { setDuration, duration, times, setTimes, onInstantSubmit, onDelete, onDeleteAll, onStop, onGetAll, laoding } = useTime()
    const [slideOverOpen, setSlideOverOpen] = useState(false)
    const [totalTime, setTotalTime] = useState('00:00:00')
    const router = useRouter()

    const onCreateSchedulerSlideOver = () => {
        setSlideOverOpen(!slideOverOpen);
    }

    const onRefreshList = () => {
        onGetAll(params.branch).then(() => {
            handleHeader();
        });
    }

    const onBackToBranches = async () => {
        router.back()
    }

    const onInstantRecord = async () => {
        setDuration(0)
        await onInstantSubmit();
    }
    const handleHeader = () => {
        let totalDuration = moment.duration(0);
        times.map(time => {
            if (time.startDateTime && time.endDateTime) {
                var duration = moment.duration(moment(time.endDateTime).diff(moment(time.startDateTime)));
                totalDuration.add(duration)
            }
        })
        setTotalTime(moment({ hour: totalDuration.hours(), minutes: totalDuration.minutes(), seconds: totalDuration.seconds() }).format('HH:mm:ss'));
    }
    useEffect(() => {
        onGetAll(params.branch).then(() => {
            handleHeader();
        });

    }, [])

    return (
        <Fragment>
            <SlideOver openState={slideOverOpen} openHandler={onCreateSchedulerSlideOver} title="Create New Time">
                <form>
                    <div className="flex flex-col">
                        {/* <div>
                            <Input
                                id="duration"
                                label="Time Duration"
                                type="text"
                                extra="mb-3"
                                onChange={setDuration}
                                value={duration}
                            />
                        </div>
                        <div>
                            <Input
                                id="startPointDateTime"
                                label="Start Point Date Time"
                                type="text"
                                extra="mb-3"
                                onChange={setStartPointDateTime}
                                value={startPointDateTime}
                            />
                        </div> */}
                        {/* <div>
                            <MediumButton type="button" text="Create" extraClasses="w-full" onClick={() => onSubmit()} roundedSize="lg">
                                <AiOutlinePlus />
                            </MediumButton>
                        </div> */}
                        <p>Working on it</p>
                    </div>
                </form>
            </SlideOver>
            <PageActionMenu actions={[
                {
                    icon: <RxPlay />,
                    text: 'Instant Record',
                    position: positions.TopRight,
                    handler: onInstantRecord,
                    classes: 'rounded-lg dark:bg-white-800 bg-indigo-600 px-3 text-white hover:bg-brand-600 active:bg-brand-700  dark:hover:bg-brand-300 dark:active:bg-brand-200  rounded-lg w-full p-2 text-xs text-base font-medium text-white transition duration-200 hover:bg-brand-600 active:bg-brand-700 dark:bg-brand-400 dark:text-white dark:hover:bg-brand-300 dark:active:bg-brand-200 '
                },
                {
                    icon: <TfiLayersAlt />,
                    text: 'Scheduler',
                    position: positions.TopRight,
                    handler: onCreateSchedulerSlideOver,
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
                    text: 'Back to branches',
                    position: positions.TopLeft,
                    handler: onBackToBranches,
                    classes: 'bg-gray-100 px-4 py-2 text-base font-medium text-slate-500 transition duration-200 hover:bg-gray-200 active:bg-gray-300 dark:bg-white/10 dark:text-white dark:hover:bg-white/20 dark:active:bg-white/30 '
                }
            ]} filter={[]} />

            <div className="flex flex-row w-full p-2 bg-stone-700 text-gray-100 rounded-lg mb-2 mt-0 text-base text-center">
                <div className="flex-1">Total Of Time Chunks: {totalTime} </div>
            </div>

            {laoding ? <LoadingSpinner width={28} height={28} /> : (
                times.length > 0 ? (
                    <Fragment>
                        <div className="grid grid-cols-4 gap-4">{(times.map((time, i) => <TimeCard key={i} time={time} onDelete={onDelete} onStop={onStop} />))}</div>
                    </Fragment>)
                    :
                    <Fragment>
                        <p className="w-full text-slate-500 dark:text-amber-50 text-center">{`Your branch is empty.Tik Tak`}</p>
                    </Fragment>
            )}
        </Fragment>
    );

}