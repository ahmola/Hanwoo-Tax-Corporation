import { NoticeItem } from "./Constants";

export async function getNotices(): Promise<NoticeItem[]> {
  const apiUrl = process.env.NEXT_PUBLIC_API_URL;
  
  if (!apiUrl) {
    console.error("API URL is not defined in .env.local");
    return [];
  }

  try {
    const res = await fetch(`${apiUrl}/notices`, {
        next: { revalidate: 60 * 60 * 24} // 24시간마다 캐시 갱신
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