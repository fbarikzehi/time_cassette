'use client'

import { Fragment, useEffect } from "react";

export default function CassetteLayout({
    children,
}: {
    children: React.ReactNode
}) {
    return (
        <Fragment>
            {/*
              //TODO: feature - add a tape indicator on top of the lists of cassettes,fragments,branches and times
             <div className="flex flex-row w-full bg-stone-700 text-gray-100 rounded-sm mb-2 mt-0 text-base text-center">
                <div className="flex-1">Select a cassette fragments</div>
            </div> */}
            {children}
        </Fragment>
    );
}