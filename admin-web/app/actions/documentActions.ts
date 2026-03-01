'use server'

import path from "path";
import fs from 'fs/promises';

export async function uploadDocument(formData: FormData) {
    const API_URL = process.env.NEXT_PUBLIC_API_URL;
    const RESOURCE_PATH = process.env.NEXT_PUBLIC_RESOURCE_PATH;
    let fileName;
    let size;

    // 1. 웹 서버에서 자체적으로 문서 저장
    try{
        const file = formData.get("file") as File;

        if (!file) {
            return {success: false, error: "파일이 없습니다."}
        }

        const buffer = Buffer.from(await file.arrayBuffer());
        fileName = `${Date.now()}-${file.name}`;
        const filePath = path.join(`${RESOURCE_PATH}`, fileName);

        await fs.mkdir(`${RESOURCE_PATH}`, {recursive: true});
        await fs.writeFile(filePath, buffer);

        size = file.size.toString();

    }catch(error){
        console.error("문서 업로드 실패", error)
        return {success: false, error: "문서 업로드 실패"}
    }

    // 2. API 서버로 문서 정보 전송
    try{
        const apiData = {
            title: formData.get("title") as string,
            category: formData.get("category") as string,
            file_url: `/api/resources/${fileName}`,
            size: size as string
        }

        console.log("apiData", apiData)

        const res = await fetch(`${API_URL}/documents`, {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify(apiData),
        });

        if (!res.ok) {
            return {success: false, error: "문서 정보 업로드 실패"}
        }

        return {
            success: res.ok,
            status: res.status
        }
    }catch(error){  
        console.error("문서 업로드 실패", error)
        return {success: false, error: "문서 정보 업로드 실패"}
    }
}