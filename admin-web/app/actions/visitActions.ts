'use server'

import fs from 'fs/promises';
import { unstable_noStore as noStore } from 'next/cache';

const VISIT_FILE = process.env.VISIT_FILE || '/usr/share/next-resources/visitor_count.json';

export async function getVisitorCount() {
  noStore();

  try {    
    console.log("Reading Visitor JSON File...")
    const content = await fs.readFile(`${VISIT_FILE}`, 'utf-8');
    const data = JSON.parse(content);

    console.log("Visit Data: ", data);

    return data.count;
  } catch (e) {
    return 0;
  }
}