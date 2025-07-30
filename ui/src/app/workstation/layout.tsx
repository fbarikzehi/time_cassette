'use client'

import { Suspense, useEffect } from "react";
import { Header,PageLoading } from "@/components";

import { NavigationEvents } from "./navigation-events";

export default function DashboardLayout({
    children,
}: {
    children: React.ReactNode
}) {
    return (
        <div className="min-h-full h-screen  dark:bg-slate-700 bg-gray-100 dark:text-white text-gray-800">
            <Header />
            <main>
                <div className="mx-auto max-w-7xl py-6 sm:px-6 lg:px-8">
                    {children}
                    <Suspense fallback={<PageLoading />}>
                        <NavigationEvents />
                    </Suspense>
                </div>
            </main>
        </div>
    );
}