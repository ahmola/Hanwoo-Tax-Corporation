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
        
        const rawData = await res.json();
        console.log("Raw Data: ", JSON.stringify(rawData, null, 2));

        const items = rawData.map((item: any) => ({
            ...item,
            date: item.created_at.split('T')[0]
        }));

        return items;
    } catch (error) {
        console.error("Document API Error:", error);
        return [];
    }
}