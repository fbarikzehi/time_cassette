'use client'

import { ExtraSmallButton, SeletcList } from "@/components";
import { positions } from "@/utils/dom.utils";
import { Fragment, useEffect } from "react";

export default function PageActionMenu(props: {
    actions: Array<{
        icon: React.ReactNode,
        text: string,
        position: positions,
        handler: React.MouseEventHandler,
        classes: string
    }>;
    filter: Array<{
        name: string,
        value: string,
        selected: boolean
    }>
}) {
    const { actions, filter, ...rest } = props;

    const onSelect = (a: any) => {
        console.log(a)
    }
    return (
        <Fragment>
            <div className="grid grid-cols-4 gap-4 mb-4">
                <div className="col-span-2">
                    {actions ? actions.filter(action => action.position == positions.TopLeft)?.map((action, key) => {
                        return (
                            <div className="h-50 mr-4 flex  flex-row justify-start" key={key}>
                                <button onClick={action.handler} type="button" className={`flex flex-row items-center text-center justify-center ${action.classes}`}>
                                    <div className="flex h-5 w-5 items-center justify-center text-lg">
                                        {action.icon ? action.icon : ''}
                                    </div>
                                    {action.text ? action.text : ''}
                                </button>
                            </div>
                        )

                    }) : ('')}
                    {filter?.length > 0 ? <SeletcList options={filter} onSelect={onSelect} /> : ''}
                </div>
                <div className="col-span-2">
                    <div className="flex flex-row  justify-end">
                        {actions ? actions.filter(action => action.position == positions.TopRight)?.map((action, key) => {
                            return (
                                <div className="h-50 ml-4 flex w-32 justify-start" key={key}>
                                    <button onClick={action.handler} type="button" className={`flex flex-row items-center text-center justify-center ${action.classes}`}>
                                        {action.text ? action.text : ''}
                                        <div className="flex h-5 w-5 items-center justify-center text-lg">
                                            {action.icon ? action.icon : ''}
                                        </div>
                                    </button>
                                </div>
                            )

                        }) : (<p>No action provided</p>)}
                    </div>
                </div>
            </div>
        </Fragment>

    );
}