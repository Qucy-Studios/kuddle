import LZString from "lz-string";
import type {UploadResult} from "$lib/types";
import {kuddle} from "$lib/wailsjs/go/models";
import Session = kuddle.Session;

export const getRecentUploads = async (): Promise<UploadResult[]> => {
    let recentsRawValue = localStorage.getItem("kddle.recents")
    recentsRawValue = recentsRawValue == null ? "[]" : LZString.decompressFromUTF16(recentsRawValue)
    return await JSON.parse(recentsRawValue) as UploadResult[]
}

export const pushRecentUpload = async (session: Session, link: string): Promise<UploadResult> => {
    let uploadResult = {
        imageLink: link,
        creationTime: new Date(),
        session: session.name
    }

    let recents = await getRecentUploads()
    recents.unshift(uploadResult)

    const recentsResult = LZString.compressToUTF16(JSON.stringify(recents))
    localStorage.setItem("kddle.recents", recentsResult)

    return uploadResult
}