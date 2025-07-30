
'use client'

import { Card, CassetteIcon, ExtraSmallButton } from "@/components";

import { CassetteModel } from '@/models/Cassette'
import { Fragment } from "react";


export default function CassetteCard(props: {
    cassette: CassetteModel;
    onFetchFragments: (id: string) => void;
    onUpdateOpen: (id: string, name: string) => void;
    onDelete: (id: string) => void;
}) {
    const { cassette, onFetchFragments, onUpdateOpen, onDelete } = props;

    return (
        <Fragment>
            <Card key={cassette.name} extra="!flex-row items-center rounded-[5px] mb-3 shadow-md">
                <div className="ml-[18px] flex h-[90px] w-full flex-row items-center">
                    <div className="rounded-full bg-lightPrimary p-3 dark:bg-navy-700">
                        <span className="flex items-center text-brand-500 dark:text-white">
                            <CassetteIcon width={80} color={`#${cassette.color}`} />
                        </span>
                    </div>
                    <div className="h-50 ml-4 flex w-7/12 flex-col justify-center">
                        <p className="font-dm text-sm font-medium text-neutral-400">
                            <span>{`${cassette.counts.fragment} Fragments`} {' | '}</span>
                            {/* <span>{`${(cassette.is_private ? 'Private' : 'Public')}`} {' | '}</span>
                            <span>{`${(cassette.status ? 'Playing' : 'Paused')}`}</span> */}
                        </p>
                        <h4 className="text-xl font-bold text-navy-700 text-slate-700">
                            {cassette.name}
                        </h4>
                    </div>
                    <div className="h-50 mr-4 flex w-full  flex-row justify-end">
                        <div className="h-50 mr-4 flex w-32  flex-row justify-end">
                            <ExtraSmallButton roundedSize="lg" text="Fragments" type="button" onClick={() => onFetchFragments(cassette.Id)} extraClasses="bg-sky-500" />
                        </div>
                        <div className="h-50 mr-4 flex w-32 flex-row justify-end">
                            <ExtraSmallButton roundedSize="lg" text="Edit" type="button" onClick={() => onUpdateOpen(cassette.Id, cassette.name)} extraClasses="bg-indigo-500" isAsync={false} />
                        </div>
                        <div className="h-50 mr-4 flex w-32 flex-row justify-end">
                            <ExtraSmallButton roundedSize="lg" text="Delete" type="button" onClick={() => onDelete(cassette.Id)} extraClasses="bg-red-500" />
                        </div>
                    </div>
                </div>
            </Card>
        </Fragment>
    )
}
