'use server'

export async function uploadNotice(formData: any) {
    const API_URL = process.env.NEXT_PUBLIC_API_URL;

    try {
        const res = await fetch(`${API_URL}/notices`, {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify(formData),
        });

        if (!res.ok) {
            return {success: false, error: "공지사항 업로드 실패"}
        }

        return {
            success: res.ok,
            status: res.status
        }
    }catch(error){
        console.error("공지사항 업로드 실패", error)
        return {success: false, error: "공지사항 업로드 실패"}
    }
}