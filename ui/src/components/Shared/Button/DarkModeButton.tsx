'use client'

import { useState, useEffect } from "react";
import { useDispatch } from "react-redux";

import { change } from "@/store/slices/themeSlice";
import { MoonIcon, SunIcon } from "@/components";

export default function DarkModeButton() {

  const dispatch = useDispatch();
  const [themeMode, setThemeMode] = useState('light')

  let themeState: ThemeStateModel = { theme: { mode: '' } }

  useEffect(() => {
    const storedTheme = window.localStorage.getItem('theme') ?? 'light';
    document.body.classList.add(storedTheme);
    document.documentElement.classList.add(storedTheme);
    setThemeMode(storedTheme)
    let theme: ThemeStateModel = { theme: { mode: storedTheme } }
    dispatch(change(theme))
  }, [dispatch, setThemeMode])

  const toggleDarkMode = () => {
    if (themeMode === 'light') {
      document.body.classList.add("dark");
      document.documentElement.classList.add("dark");
      window.localStorage.setItem('theme', 'dark');
      setThemeMode('dark')
      themeState.theme.mode = 'dark';
      dispatch(change(themeState))
    } else {
      document.body.classList.remove("dark");
      document.documentElement.classList.remove("dark");
      window.localStorage.setItem('theme', 'light');
      setThemeMode('light')
      dispatch(change(themeState))
    }
  };



  return (
    <button onClick={toggleDarkMode} className={`fixed flex-space-x-1 inline-block bg-slate-200 rounded-full mb-6 ml-6 w-15 p-5 bottom-0 left-3 ${themeMode === 'light' ? 'bg-slate-800 text-white' : 'bg-slate-200 text-slate'}`}>
      {(themeMode === 'light' ? <MoonIcon /> : <SunIcon />)}
    </button>
  );
}
