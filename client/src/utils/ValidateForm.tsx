import type { FormDataType } from "../types/Alert"

export const checkEmailValidity = (email: string) => {
    // if email is not the correct email type return message for the user
    const regx = /^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$/
    if (!regx.test(email)) {
        return false
    }
    return true
}




export const validateFormData = (data: FormDataType) => {
    // check if email is valid
    if (!checkEmailValidity(data.email)) {
        return false
    }

    return true;
}


