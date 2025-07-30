import '../assets/styles/globals.css'
import DarkModeButton from "@/components/Shared/Button/DarkModeButton";
import { StoreProvider } from "@/store/storeProvider";
import { Toaster } from 'react-hot-toast';

import type { Metadata } from 'next'
import { Inter } from 'next/font/google'

const inter = Inter({ subsets: ['latin'] })

export const metadata: Metadata = {
  title: 'Time Cassette',
  description: 'Time Cassette',
}


export default function RootLayout({
  children,
}: {
  children: React.ReactNode
}) {
  return (
    <html lang="en" className='min-h-full h-screen  dark:bg-slate-700 bg-gray-100 dark:text-white text-gray-800'>
      <body className='overflow-x-hidden dark:bg-slate-700'>
        {/* <AuthProvider> */}
          <StoreProvider>
            <Toaster
              position="top-center"
              reverseOrder={false}
              gutter={8}
              containerClassName=""
              containerStyle={{}}
              toastOptions={{
                className: '',
                duration: 3000,
                style: {
                  background: '#363636',
                  color: '#fff',
                  minWidth: 400
                },
                success: {
                  duration: 2500,
                },
                error: {
                  duration: 3000,
                },
                loading: {

                }

              }}
            />
            {children}
            <DarkModeButton />
          </StoreProvider>
        {/* </AuthProvider> */}
      </body>
    </html>
  )
}
