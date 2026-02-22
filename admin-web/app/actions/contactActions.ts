'use server'

export async function getContacts() {
  const API_URL = process.env.NEXT_PUBLIC_API_URL;

  try {
    const res = await fetch(`${API_URL}/contacts/all`, {
      cache: 'no-store', // 매번 새로운 데이터를 가져오도록 설정
    });

    if (!res.ok) throw new Error("데이터를 가져오는데 실패했습니다.");

    const data = await res.json();

    return data.map((item:any) => {
      const rawDate = item.created_at;
      
      return {
        ...item,
        created_at: rawDate ? new Date(rawDate).toLocaleDateString('ko-KR', {
            year: 'numeric',
            month: 'long',
            day: 'numeric',
            hour: '2-digit',
            minute: '2-digit'
        })
        : "날짜 정보 없음"
      }
    })
  } catch (err) {
    console.error("Server Action 에러:", err);
    return [];
  }
}