'use server'

import { NoticeItem } from "@/features/notice/Constants";

export async function getNotices(): Promise<NoticeItem[]> {
  const apiUrl = process.env.NEXT_PUBLIC_API_URL;
  
  if (!apiUrl) {
    console.error("API URL is not defined in .env.local");
    return [];
  }

  try {
    const res = await fetch(`${apiUrl}/notices`, {
        cache: 'no-store'
    });

    if (!res.ok) {
      throw new Error(`Failed to fetch notices: ${res.status}`);
    }

    return res.json();

  } catch (error) {
    console.error("Notice API Error:", error);
    return [];
  }
}