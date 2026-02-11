import { DocumentItem } from "./Constants";

export async function getDocuments(): Promise<DocumentItem[]> {
    const apiUrl = process.env.NEXT_PUBLIC_API_URL;

    if (!apiUrl) {
        console.error("API URL is undefined in .env.local");
        return [];
    }

    try {
        const res = await fetch(`${apiUrl}/documents`, {
            next : {revalidate: 60*60*24} // 24시간마다 갱신
        });

        if (!res.ok) {
            throw new Error(`Failed to fetch documents: ${res.status}`);
        }

        return res.json();
    } catch (error) {
        console.error("Document API Error:", error);
        return [];
    }
}