'use client'

import { useEffect, useState } from "react";
import { HiOutlineTrash, HiStopCircle } from "react-icons/hi2";
import { BiStopCircle } from "react-icons/bi";
import moment from 'moment';
import { Card, FragmentIcon, ExtraSmallButton } from "@/components";
import { TimeModel } from '@/models/Time'
import { IoChevronDownOutline } from "react-icons/io5";

export default function TimeCard(props: {
    time: TimeModel;
    onDelete: Function,
    onStop: Function,
}) {
    const { time, onDelete, onStop } = props;

    const [elapsedSeconds, setElapsedSeconds] = useState(1);
    const [currentTime, setCurrentTime] = useState('00:00:00');

    const onUpdate = () => {
        setInterval(() => {
            setElapsedSeconds(elapsedSeconds + 1);
            let currentElapsedSeconds = (moment().unix() - moment(time.startDateTime).unix()) + elapsedSeconds;
            setCurrentTime(moment.unix(currentElapsedSeconds).format('HH:mm:ss'));
        }, 1000)
    }
    useEffect(() => {
        if (time.startDateTime && !time.endDateTime) {
            setCurrentTime(moment.unix(moment().unix() - moment(time.startDateTime).unix()).format('HH:mm:ss'));
            onUpdate();
        }
        else if (time.startDateTime && time.endDateTime) {
            var duration = moment.duration(moment(time.endDateTime).diff(moment(time.startDateTime)));
            setCurrentTime(moment({ hour: duration.hours(), minutes: duration.minutes(), seconds: duration.seconds() }).format('HH:mm:ss'));
        }
    }, [])
    return (
        <div className={`!z-5 flex-col relative flex rounded-[5px] ${time.startDateTime && !time.endDateTime ? 'gradient-animated-bg' : ''} bg-white bg-clip-border shadow-md shadow-shadow-500 !flex-row flex-grow items-center rounded-[20px]`}>
            <div className="flex flex-col">
                <div className="flex flex-row p-2">
                    <button className="dark:hover:bg-brand-300 dark:active:bg-brand-200 text-base font-medium text-white transition duration-200 dark:text-white" id="popover-trigger-:R29d8mH1:" aria-haspopup="dialog" aria-expanded="true" aria-controls="popover-content-:R29d8mH1:">
                        <IoChevronDownOutline />
                    </button>
                </div>
                <div className="flex flex-row p-2">
                    <div className="ml-[8px] flex h-[70px] w-auto flex-row items-center">
                        <div className="rounded-full bg-lightPrimary">
                            <span className="flex items-center text-brand-500 dark:text-stone-700">
                                {time.startDateTime && !time.endDateTime ? (<button type="button" onClick={() => onStop(time.Id)} className="rounded-xl bg-gray-100 px-5 py-3 text-3xl text-navy-700 transition duration-10 hover:bg-gray-200 active:bg-gray-300">
                                    <BiStopCircle />
                                </button>) : (
                                    <button type="button" onClick={() => onDelete(time.Id)} className="rounded-xl bg-red-100 px-5 py-3 text-3xl text-navy-700 transition duration-10 hover:bg-gray-200 active:bg-gray-300">
                                        <HiOutlineTrash />
                                    </button>
                                )}
                            </span>
                        </div>
                    </div>
                    <div className="h-50 ml-4 flex w-auto flex-col justify-center">
                        <h4 className="text-xl font-bold dark:text-stone-700">
                            {currentTime}
                        </h4>
                        <p className="font-dm text-sm font-medium text-gray-300">{time.name}</p>
                    </div>
                </div>
            </div>
        </div>
    )
}






