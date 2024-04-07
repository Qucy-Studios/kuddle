import {writable} from "svelte/store";

export const currentSessionName = writable(null as string | null)
export const sessionNames = writable(null as string[] | null)
export const encryptionKey = writable(null as string | null)
export const hasRegistered = writable(null as boolean | null)
export const isOffline = writable(false)
export const wasAutoLoggedOut = writable(false)