'use client'

import { ChangeEvent, Fragment, useState, useEffect } from "react";
import { useSelector } from "react-redux";
import Image from 'next/image'
import Link from 'next/link'
import { useRouter } from 'next/navigation'
import { InitLoginModel, LoginModel } from "@/models/LoginModel";
import { AuthService } from "@/services/AuthService";
import { useSession } from '@/hooks/useSession'



export default function Login() {
  const appState = useSelector((state: storeModel) => state.appState);

  const [model, setModel] = useState<LoginModel>(InitLoginModel);
  const [loading, setLoading] = useState(false);
  const router = useRouter()
  const { setSession } = useSession()

  const handleChange = (event: ChangeEvent<HTMLInputElement>) => {
    setModel(currentModel => ({
      ...currentModel,
      [event.target.name]: event.target.value,
    }));
  };

  const onSubmit = async () => {
    setLoading(true);
    AuthService().login(model).then((response) => {
      if (response.meta.result) {
        setSession(response.data)
        router.push(appState.activeMenu as string ?? "/workstation")
      }
      setLoading(false);
    })
  }
  return (
    <>
      <div className="flex min-h-full flex-1 flex-col justify-center px-12 py-12 lg:px-8 bg-slate-50 dark:bg-slate-900 rounded-md w-96 shadow-lg">
        {!loading ? (<Fragment>
          <div className="sm:mx-auto sm:w-full sm:max-w-sm">
            <Image
              src="/login.png"
              alt="time cassette Logo"
              className="dark:invert mx-auto h-13 w-auto"
              width={70}
              height={70}
              priority
            />
            <h2 className="mt-10 text-center text-2xl font-bold leading-9 tracking-tight text-gray-900 dark:text-white">
              Sign in to your account
            </h2>
          </div>
          <div className="mt-10 sm:mx-auto sm:w-full sm:max-w-sm">
            <form className="space-y-6" action="#" method="POST">

              <div>
                <div className="mt-2 ">
                  <label htmlFor="email" className="block text-sm font-medium leading-6 text-gray-900 dark:text-white">
                    Email address
                  </label>
                  <input
                    value={model.email}
                    onChange={(event) => handleChange(event)}
                    id="email"
                    name="email"
                    type="email"
                    autoComplete="email"
                    required
                    className="block w-full rounded-md border-0 py-1.5 px-2 dark:text-white text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-1 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
                  />
                </div>
              </div>

              <div>
                <div className="mt-2">
                  <div className="flex items-center justify-between">
                    <label htmlFor="password" className="block text-sm font-medium leading-6 text-gray-900 dark:text-white">
                      Password
                    </label>
                    <div className="text-sm">
                      <a href="#" className="font-semibold text-indigo-600 hover:text-indigo-500 dark:text-indigo-200">
                        Forgot password?
                      </a>
                    </div>
                  </div>
                  <input
                    value={model.password}
                    onChange={(event) => handleChange(event)}
                    id="password"
                    name="password"
                    type="password"
                    autoComplete="current-password"
                    required
                    className="block w-full rounded-md border-0 py-1.5  px-2 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-1 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
                  />
                </div>
              </div>
              <div>
                <button
                  onClick={onSubmit}
                  disabled={loading}
                  type="button"
                  className="flex w-full justify-center rounded-md bg-indigo-600 px-3 py-1.5 text-sm font-semibold leading-6 text-white  dark:bg-slate-700 dark:text-white  shadow-sm hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600"
                >
                  Sign in
                </button>
              </div>
            </form>

            <p className="mt-10 text-center text-sm text-gray-500">
              Do not have an account? {' '}
              <Link href="/auth/signup" className="font-semibold leading-6 text-indigo-600 hover:text-indigo-500 dark:text-indigo-200">
                Signup
              </Link>
            </p>
          </div>
        </Fragment>) : (<Fragment>
          <p className="text-center">Signing in...</p>
        </Fragment>)}
      </div>
    </>

  )
}
