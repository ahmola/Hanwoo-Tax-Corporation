'use client'

import { Download } from "lucide-react";

export default function DownloadButton({ fileUrl, title }: { fileUrl: string, title: string }) {
    const downloadApiUrl = `/api/download?url=${encodeURIComponent(fileUrl)}`;
  
    return (
    <a 
      download = {fileUrl.split('/').pop()}
      href={downloadApiUrl}
      onClick={(e) => {
        alert("Download")
      }} // 클릭 시 로그 출력
      className="p-2 text-slate-400 hover:text-blue-700 hover:bg-slate-100 rounded-full transition ml-2 flex-shrink-0"
      title="다운로드"
    >
      <Download className="w-5 h-5" />
    </a>
  );
}