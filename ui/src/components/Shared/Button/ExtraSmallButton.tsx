'use client'
import { LoadingSpinner } from "@/components";
import { useState } from "react";


export default function ExtraSmallButton(props: {
    text: string;
    type: "button" | "submit";
    onClick: (clickEventParams: Function) => any;
    clickEventParams?: any;
    isAsync?: boolean;
    roundedSize: string;
    extraClasses?: string;
}) {
    const { text, type, onClick, isAsync = true, clickEventParams, roundedSize, extraClasses, ...rest } = props;

    const [loading, setLoading] = useState(false)

    const handleClick = async () => {
        setLoading(true);
        await onClick(clickEventParams).then(() => {
            setLoading(false);
        });
    }

    return (
        <>
            <button onClick={() => isAsync ? handleClick() : onClick(clickEventParams)} type={type} className={`rounded-${roundedSize} w-full p-2 text-xs font-medium text-white transition duration-200 hover:bg-brand-600 active:bg-brand-700 dark:bg-brand-400 dark:text-white dark:hover:bg-brand-300 dark:active:bg-brand-200 ${extraClasses}`}>
                {loading ? <LoadingSpinner width={16} height={16} /> : text}
            </button>
        </>
    );
}