'use client'

import { useEffect, useState } from "react";
import Image from "next/image"
import { useSelector } from "react-redux";

export default function MainLogo() {
  const themeMode = useSelector<ThemeStateModel>(state => state.theme.mode);
  const [mode, setMode] = useState(themeMode)

  useEffect(() => {
    setMode(window.localStorage.getItem('theme') ? window.localStorage.getItem('theme') : 'light')
  }, [])

  return (
    <Image
      className="h-8 w-8"
      src={`/cassettes/for-${mode}-cassette.svg`}
      width={35}
      height={35}
      alt="Time Cassette"
    />
  );
}
