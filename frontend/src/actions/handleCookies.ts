"use server"

import { cookies } from "next/headers"


export const storeCookie = async (name: string, value: string) => {
    cookies().set(name, value, {maxAge: 86400})
}

export const getCookie = async (name: string) => {
    return cookies().get(name)
}


export const deleteCookie = async (name: string) => {
    cookies().delete(name)
}