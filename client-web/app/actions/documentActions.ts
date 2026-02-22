'use server'

import { DocumentItem } from "@/features/documents/Constants";

export async function getDocuments(): Promise<DocumentItem[]> {
    const apiUrl = process.env.NEXT_PUBLIC_API_URL;

    if (!apiUrl) {
        console.error("API URL is undefined in .env.local");
        return [];
    }

    try {
        const res = await fetch(`${apiUrl}/documents`, {
            cache: 'no-store'
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