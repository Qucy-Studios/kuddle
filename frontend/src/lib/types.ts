export type AlertError = { title: string, description: string }
export type Session = {
    name: string,
    accessKeyId: string,
    endpoint: string,
    secretAccessKey: string,
    token: string,
    useSsl: boolean,
    bucket: string,
    domain: string
}
export type UploadResult = {
    imageLink: string,
    creationTime: Date,
    session: string
}
export type SessionSettings = {
    urlLength: number,
    folder: string,
}