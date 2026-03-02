'use server'

import { DocumentItem } from "@/features/documents/Constants";
import { NextResponse } from "next/server";

export async function getDocumentsInfo(): Promise<DocumentItem[]> {
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
        
        const data = await res.json();

        console.log("Data: ", JSON.stringify(data, null, 2));

        return data;
    } catch (error) {
        console.error("Document API Error:", error);
        return [];
    }
}