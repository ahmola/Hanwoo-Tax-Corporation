'use server'

export async function uploadDocument(formData: FormData) {
    const API_URL = process.env.NEXT_PUBLIC_API_URL;

    try{
        const res = await fetch(`${API_URL}/documents`, {
            method: "POST",
            body: formData, // JSON.stringify 하지 않음!
        });

        if (!res.ok) {
            return {success: false, error: "문서 업로드 실패"}
        }

        return {
            success: res.ok,
            status: res.status
        }
    }catch(error){  
        console.error("문서 업로드 실패", error)
        return {success: false, error: "문서 업로드 실패"}
    }
}