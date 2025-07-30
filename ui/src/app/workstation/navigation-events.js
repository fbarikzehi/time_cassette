'use client'

import { useEffect } from 'react'
import { usePathname, useSearchParams } from 'next/navigation'
import { useSession } from "@/hooks/useSession";
import { useRouter } from 'next/navigation'

export function NavigationEvents() {
      const pathname = usePathname()
      const searchParams = useSearchParams()

    const { isValidSession } = useSession()
    const router = useRouter()
    useEffect(() => {
        if (!isValidSession())
            router.push("/auth/signout")
    }, [pathname,searchParams])

    return null
}