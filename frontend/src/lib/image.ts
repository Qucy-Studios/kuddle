import type {AlertError} from "$lib/types";

export const isValidBase64Image = (data: string | null): AlertError | null => {
    if (data === "" || data == null) {
        return {
            title: "Unknown clipboard data.",
            description: "We weren't unable to determine what was sent through the clipboard."
        }
    }

    const contentType = data.split(';')[0].replace("data:", "")
    if (!contentType.startsWith("image/")) {
        return {
            title: "Invalid file format.",
            description: "Kuddles only accept image files, if this is an image file, then please report this as a bug."
        }
    }

    return null
}

export const extractImageSrc = (html: string) => {
    const regex = /<img[^>]+src\s*=\s*['"]([^'"]+)['"][^>]*>/g;
    const matches = [];
    let match;

    // Loop through matches
    while ((match = regex.exec(html)) !== null) {
        // match[1] contains the value of the src attribute
        matches.push(match[1]);
    }

    return matches;
}