// 상담 폼
'use client'

import { submitConsultation } from "@/app/actions";
import { useState } from "react";

export default function ContactForm() {
    const [isSubmitted, setIsSubmitted] = useState(false);
    const [isPending, setIsPending] = useState(false);

    async function handleSubmit(formData: FormData) {
        setIsPending(true);
        const result = await submitConsultation(formData);
        setIsPending(false);

        if (result.success) {
            setIsSubmitted(true);
        } else {
            alert("알 수 없는 오류가 발생했습니다. 잠시 후 다시 시도해주세요.");
        }

    }

    if (isSubmitted) {
        return (
            <div className="container mx-auto px-4 py-20 flex justify-center">
                <div className="max-w-md w-full p-10 bg-green-50 text-green-700 rounded-2xl text-center border border-green-200">
                    <h3 className="text-xl font-bold mb-2">상담 신청 완료!</h3>
                    <p>작성하신 내용이 성공적으로 전달되었습니다.<br/>담당자가 확인 후 곧 연락드리겠습니다.</p>
                </div>
            </div>
        );
    }

    return (
        <section id="contact" className="bg-slate-50 py-20 px-4">
            <div className="container mx-auto text-center">
            <h2 className="text-3xl font-bold mb-8 text-slate-900">상담 신청</h2>
            <form action={handleSubmit} className="space-y-4 max-w-md mx-auto bg-white p-6 rounded-xl shadow-lg border">
                <h3 className="text-lg font-bold text-slate-800 mb-2">전문가 무료 상담</h3>
                <div>
                    <label className="block text-sm font-medium text-slate-600 mb-1">성함</label>
                    <input name="name" required className="input-primary" placeholder="홍길동" />
                </div>
                <div>
                    <label className="block text-sm font-medium text-slate-600 mb-1">연락처</label>
                    <input name="phone" required className="input-primary" placeholder="010-0000-0000" />
                </div>
                <div>
                    <label className="block text-sm font-medium text-slate-600 mb-1">문의 내용</label>
                    <textarea name="message" className="input-primary" placeholder="조세 문의 내용" />
                </div>
                <button 
                    type="submit" 
                    disabled={isPending}
                    className={`w-full py-4 text-white font-bold rounded-xl transition shadow-md ${isPending ? 'bg-slate-400 cursor-not-allowed' : 'bg-blue-700 hover:bg-blue-800'}`}
                >
                    {isPending ? "전송 중..." : "상담 신청하기"}
                </button>
            </form>
        </div>
      </section>
    );
}