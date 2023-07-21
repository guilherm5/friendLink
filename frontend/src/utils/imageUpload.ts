type ImageUploadReturn = {error: boolean, image: File | null, urlPreview: string | null}

const setPreview = (file: File, callback: Function) => {
    const reader = new FileReader()
    reader.readAsDataURL(file)
    reader.onloadend = () => {
        callback(reader.result)
    };
};

const imageUpload = (e: Event, callback: Function | undefined = undefined) => {
    if (!e){return {error: true, image: null, urlPreview: null}}
    const target = e.target as HTMLInputElement;
    if(!target.files){return {error: true, image: null, urlPreview: null}}
    const image = target.files[0];
    return setPreview(image, (url: string):ImageUploadReturn => {
        if (callback){
            return callback({error: false, image, urlPreview: url})
        }
        return {error: false, image, urlPreview: url}
    })
}
export {
    setPreview,
    imageUpload
}

export type {
    ImageUploadReturn
}