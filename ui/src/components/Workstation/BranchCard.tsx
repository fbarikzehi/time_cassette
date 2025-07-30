
'use client'

import { Card, BranchIcon, ExtraSmallButton } from "@/components";
import { BranchModel } from "@/models/Branch";

export default function BranchCard(props: {
    branch: BranchModel;
    onFetchTimes: Function,
    onUpdateOpen: (id: string, name: string) => void;
    onDeleteRequest: Function,
}) {
    const { branch, onFetchTimes, onDeleteRequest,onUpdateOpen } = props;

    return (
        <Card key={branch.name} extra="!flex-row items-center rounded-[5px] mb-3 shadow-md">
            <div className="ml-[18px] flex h-[90px] w-full flex-row items-center">
                <div className="rounded-full bg-lightPrimary p-3 dark:bg-navy-700">
                    <span className="flex items-center text-brand-500 dark:text-white">
                        <BranchIcon width={80} color={`#${branch.color}`} />
                    </span>
                </div>
                <div className="h-50 ml-4 flex w-7/12 flex-col justify-center">
                    <p className="font-dm text-sm font-medium text-neutral-400">
                        <span>{`${branch.counts.time} Times`}</span>
                    </p>
                    <h4 className="text-xl font-bold text-navy-700 text-slate-700">
                        {branch.name}
                    </h4>
                </div>
                <div className="h-50 mr-4 flex w-full  flex-row justify-end">
                    <div className="h-50 mr-4 flex w-32  flex-row justify-end">
                        <ExtraSmallButton roundedSize="lg" text="Times" type="button" onClick={() => onFetchTimes(branch.Id)} extraClasses="bg-sky-500" />
                    </div>
                    <div className="h-50 mr-4 flex w-32 flex-row justify-end">
                        <ExtraSmallButton roundedSize="lg" text="Edit" type="button" onClick={() => onUpdateOpen(branch.Id, branch.name)} extraClasses="bg-indigo-500" isAsync={false} />
                    </div>
                    <div className="h-50 mr-4 flex w-32 flex-row justify-end">
                        <ExtraSmallButton roundedSize="lg" text="Delete" type="button" onClick={() => onDeleteRequest(branch.Id)} extraClasses="bg-red-500" />
                    </div>
                </div>
            </div>
        </Card>
    )
}






