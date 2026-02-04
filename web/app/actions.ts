// 임시 백엔드
'use server'

export async function submitConsultation(formData: FormData) {
  const data = {
    name: formData.get('name'),
    phone: formData.get('phone'),
    message: formData.get('message'),
  };

  // TODO: 나중에 여기서 Go/Spring 백엔드 API 호출
  // await fetch('http://backend-service/api/consultation', { ... })
  
  return { success: true, message: "상담이 접수되었습니다." };
}