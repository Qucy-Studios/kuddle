import {createEventDispatcher} from "svelte";
import type {AlertError} from "$lib/types";

export let validateNotEmpty = (...strs: any[]): boolean => {
    for (let str of strs) {
        if (str == "" || str == null) return false
    }
    return true
}

export let useSetupDispatchEvents = (validation: (() => Promise<AlertError | undefined>) | undefined) => {
    const dispatcher = createEventDispatcher()
    const toggleShowSecrets = () => dispatcher('toggle')
    const nextStage = async () => {
        if (validation != null) {
            let err = await validation()
            if (err != null) {
                dispatcher('validation', err)
                return
            }
        }

        dispatcher('next')
    }

    return {
        dispatcher,
        toggleShowSecrets,
        nextStage
    }
}