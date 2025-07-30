'use client'

import { useEffect } from "react";
import { useSession } from '@/hooks/useSession'
import { useRouter } from 'next/navigation'

export default function Signout() {
    const { clearSession } = useSession()
    const router = useRouter()

    useEffect(() => {
        clearSession()
        router.push("/auth/signin")
    }, [clearSession,router])

    return (
        <div className="flex min-h-full flex-1 flex-col justify-center px-12 py-12 lg:px-8 bg-slate-50 dark:bg-slate-900 rounded-md w-96 shadow-lg">
            <p className="text-center">Signing out...</p>
        </div>
    )
}