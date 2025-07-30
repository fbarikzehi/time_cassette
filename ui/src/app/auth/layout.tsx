import type { Metadata } from 'next'
import Image from 'next/image'
import Link from "next/link";
import VerifyBadgeIcon from '@/components/Shared/Icon/VerifyBadgeIcon'
import { MainLogo } from "@/components";


export const metadata: Metadata = {
	title: 'Welcome on board',
	description: 'Time Cassette',
}

export default function AuthLayout({
	children,
}: {
	children: React.ReactNode;
}) {
	return (
		<main className="flex min-h-screen flex-col items-center justify-between p-24">
			<div className="z-10 max-w-5xl w-full items-center justify-between font-mono text-sm lg:flex">
				<div className="fixed bottom-0 left-0 flex h-48 w-full items-end justify-center bg-gradient-to-t from-white via-white dark:from-black dark:via-black lg:static lg:h-auto lg:w-auto lg:bg-none">
					<a
						className="pointer-events-none flex place-items-center gap-2 p-8 lg:pointer-events-auto lg:p-0 dark:text-white"
						href="https://vercel.com?utm_source=create-next-app&utm_medium=appdir-template&utm_campaign=create-next-app"
						target="_blank"
						rel="noopener noreferrer"
					>
						<MainLogo />
						{' '} Time Cassette
					</a>
				</div>
				<Link replace prefetch={true} href="/auth/signup" className="rounded-md flex bg-indigo-600 px-5 py-2 text-sm font-semibold text-white dark:bg-slate-800 dark:text-white shadow-sm hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600">
					<VerifyBadgeIcon className="mr-1.5" />
					Signup
				</Link>
			</div>

			<div className="relative flex place-items-center">
				{children}
			</div>

			<div className="mb-32 grid text-center lg:max-w-5xl lg:w-full lg:mb-0 lg:grid-cols-4 lg:text-left">
			</div>
		</main>
	);
}
