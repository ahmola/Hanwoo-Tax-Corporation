'use client'

import { uploadNotice } from "@/app/actions/noticeActions";
import { useState } from "react";

export default function NoticeWriteForm() {
  const [formData, setFormData] = useState({
    title: "",
    category: "공지", // 기본값
    content: "",
    is_important: false,
  });

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();

    // API 서버로 POST 요청
    console.log(formData)
    const res = await uploadNotice(formData);

    if (res.success) {
      alert("공지사항이 등록되었습니다.");
      // 입력창 초기화 or 목록으로 이동 로직
    } else {
      alert("등록 실패");
    }
  };

  return (
    <form onSubmit={handleSubmit} className="max-w-2xl bg-white p-8 rounded-xl border space-y-6">
      <h2 className="table-head-name">공지사항 등록</h2>
      
      {/* 카테고리 & 중요 체크 */}
      <div className="flex gap-4">
        <select 
          className="border p-2 rounded text-black"
          value={formData.category}
          onChange={e => setFormData({...formData, category: e.target.value})}
        >
          <option value="공지">공지</option>
          <option value="세무일정">세무일정</option>
          <option value="뉴스">뉴스</option>
        </select>
        <label className="flex items-center gap-2 text-red-500">
          <input 
            type="checkbox" 
            checked={formData.is_important}
            onChange={e => setFormData({...formData, is_important: e.target.checked})}
          />
          중요 공지
        </label>
      </div>

      <input 
        className="table-input" 
        placeholder="제목을 입력하세요"
        value={formData.title}
        onChange={e => setFormData({...formData, title: e.target.value})}
      />
      
      <textarea 
        className="table-input h-64" 
        placeholder="내용을 입력하세요"
        value={formData.content}
        onChange={e => setFormData({...formData, content: e.target.value})}
      />

      <button className="table-button">
        등록하기
      </button>
    </form>
  );
}