'use client'

import { useState } from "react";

export default function DocumentUploadForm() {
  const [title, setTitle] = useState("");
  const [file, setFile] = useState<File | null>(null);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!file) return alert("파일을 선택해주세요.");

    const formData = new FormData();
    formData.append("title", title);
    formData.append("file", file); // 백엔드에서 받을 필드명 'file'

    // API 서버로 파일 전송
    const res = await fetch(`${process.env.NEXT_PUBLIC_API_URL}/documents`, {
      method: "POST",
      body: formData, // JSON.stringify 하지 않음!
    });

    if (res.ok) {
      alert("업로드 완료");
    } else {
      alert("업로드 실패");
    }
  };

  return (
    <form onSubmit={handleSubmit} className="max-w-2xl bg-white p-8 rounded-xl border space-y-6">
      <h2 className="table-head-name">자료실 업로드</h2>
      
      <input 
        className="table-input" 
        placeholder="자료 제목"
        value={title}
        onChange={e => setTitle(e.target.value)}
      />

      <input 
        type="file"
        className="table-input"
        onChange={e => setFile(e.target.files?.[0] || null)}
      />

      <button className="table-button">
        업로드하기
      </button>
    </form>
  );
}
