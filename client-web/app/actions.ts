'use server'

import { Resend } from "resend";

const resend = new Resend(process.env.NEXT_PUBLIC_RESEND_API_KEY);

export async function submitConsultation(formData: FormData) {
  const name = formData.get('name') as string;
  const phone = formData.get('phone') as string;
  const message = formData.get('message') as string;

  // send email with Resend
  try {
    // 1. 관리자에게 알림 메일 발송
    await resend.emails.send({
      from: 'onboarding@resend.dev', // 나중에 도메인으로 변경 예정
      to: process.env.NEXT_PUBLIC_ADMIN_EMAIL as string,  // 상담 내용을 받을 실제 이메일 주소
      subject: `[신규 상담 신청] ${name} 고객님`,
      html: `
        <h2>새로운 상담 신청이 접수되었습니다.</h2>
        <p><strong>성함:</strong> ${name}</p>
        <p><strong>연락처:</strong> ${phone}</p>
        <p><strong>문의 내용:</strong></p>
        <div style="padding: 15px; background-color: #f5f5f5; border-radius: 5px;">
          ${message.replace(/\n/g, '<br>') || '내용 없음'}
        </div>
        <p style="color: #888; font-size: 12px; mt: 20px;">본 메일은 웹사이트 폼을 통해 자동 발송되었습니다.</p>
      `,
    });

  }catch (error) {
    console.error("메일 발송 실패", error);
    return { success: false, error: "상담 신청 실패"}
  }

  // 2. 접수 사항을 서버로 전달
  try{
    await fetch(`${process.env.NEXT_PUBLIC_API_URL}/contacts`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          "name": name,
          "phone": phone,
          "message": message,
        }),
      }
    )
  }catch (error) {
    console.error("접수 내역 서버로 업로드 실패", error);
    return {success: false, error: "업로드 실패"}
  }
  
  return { success: true, message: "상담이 접수되었습니다." };
}