// 상담 폼
'use client'

import { submitConsultation } from "@/app/actions";
import { useState } from "react";

export default function ContactForm() {
    const [isSubmitted, setIsSubmitted] = useState(false);

    async function handleSubmit(formData: FormData) {
        await submitConsultation(formData);
        setIsSubmitted(true);
    }

    if (isSubmitted) {
        return <div className="p-6 bg-green-50 text-green-700 rounded-lg text-center">
            상담 신청이 완료되었습니다. 곧 연락드리겠습니다.
            </div>;
    }

    return (
        <section id="contact" className="bg-slate-50 py-20 px-4">
            <div className="container mx-auto text-center">
            <h2 className="text-3xl font-bold mb-8 text-slate-900">상담 신청</h2>
            <form action={handleSubmit} className="space-y-4 max-w-md mx-auto bg-white p-6 rounded-xl shadow-lg border">
                <h3 className="text-lg font-bold text-slate-800 mb-2">전문가 무료 상담</h3>
                <div>
                    <label className="block text-sm font-medium text-slate-600 mb-1">성함</label>
                    <input name="name" required className="w-full p-2 border rounded-md focus:ring-2 focus:ring-blue-500 outline-none" placeholder="홍길동" />
                </div>
                <div>
                    <label className="block text-sm font-medium text-slate-600 mb-1">연락처</label>
                    <input name="phone" required className="w-full p-2 border rounded-md focus:ring-2 focus:ring-blue-500 outline-none" placeholder="010-0000-0000" />
                </div>
                <div>
                    <label className="block text-sm font-medium text-slate-600 mb-1">문의 내용</label>
                    <textarea name="message" className="w-full p-2 border rounded-md focus:ring-2 focus:ring-blue-500 outline-none h-24" placeholder="법인 설립 관련 문의드립니다." />
                </div>
                <button type="submit" className="w-full py-3 bg-blue-700 text-white font-bold rounded-md hover:bg-blue-800 transition">
                    상담 신청하기
                </button>
            </form>
        </div>
      </section>
    );
}