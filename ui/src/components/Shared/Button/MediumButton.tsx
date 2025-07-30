'use client'

import { LoadingSpinner } from "@/components";
import { Fragment, useState } from "react";


export default function MediumButton(props: {
    text: string;
    type: 'submit' | 'reset' | 'button' | undefined;
    onClick: (clickEventParams: any) => Promise<boolean | undefined | void>;
    clickEventParams?: any;
    roundedSize: string;
    extraClasses?: string;
    children?: JSX.Element | any[];
    extra?: string;
}) {
    const { text, children, type, onClick, clickEventParams, roundedSize, extraClasses, ...rest } = props;

    const [loading, setLoading] = useState(false)

    const handleClick = async () => {
        setLoading(true);
        await onClick(clickEventParams).then(() => {
            setLoading(false);
        });
    }
    return (
        <>
            <button onClick={() => handleClick()} type={type} className={`rounded-${roundedSize} flex flex-row items-center justify-center rounded-xl bg-green-700 dark:bg-zinc-700 hover:bg-brand-600 dark:hover:bg-brand-300 dark:active:bg-brand-200 active:bg-brand-700 dark:bg-brand-400 px-12 py-3 text-base font-medium text-white transition duration-200 ${extraClasses}`}>
                {loading ? <LoadingSpinner width={23} height={23} /> :
                    <Fragment>
                        {children}
                        {text}
                    </Fragment>}
            </button>
        </>
    );
}
