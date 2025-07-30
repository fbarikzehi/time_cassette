import { usePathname } from 'next/navigation'
import Link from "next/link"

import { classNames } from '@/utils/dom.utils'
import { MenuItems } from '@/utils/app.utils'

export default function Menu() {
    const pathname = usePathname()

    return (
        <>
            <div className="ml-10 flex items-baseline space-x-4">
                {MenuItems.map((link) => {
                    let pathNames = pathname.split('/')
                    let linkHrefs = link.href.split('/')
                    const currentActiveNavigation = linkHrefs[linkHrefs.length - 1] === pathNames[pathNames.length - 1]
                    return (
                        <Link replace
                            key={link.name}
                            href={link.href}
                            className={classNames(
                                currentActiveNavigation
                                    ? 'dark:bg-gray-500 bg-slate-900 text-gray-300'
                                    : 'text-gray-800 hover:bg-gray-700 hover:text-white',
                                'rounded-md px-3 py-2 text-sm font-medium dark:text-white'
                            )}
                            aria-current={currentActiveNavigation ? 'page' : undefined}>
                            {link.name}
                        </Link>
                    )
                })}
            </div>
        </>
    )

}