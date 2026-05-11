import { prisma } from '@/lib/prisma'
import { NextResponse } from 'next/server'

export const dynamic = 'force-dynamic'

export async function GET() {
  try {
    const submissions = await prisma.submission.findMany({
      where: { approved: true },
      orderBy: { createdAt: 'desc' },
      select: {
        contributorId: true,
        htmlContent: true,
        createdAt: true,
        percentage: true,
      },
    })

    // Group by contributorId and only keep the latest
    const latestPerContributor = new Map<string, any>()
    for (const sub of submissions) {
      if (!latestPerContributor.has(sub.contributorId)) {
        latestPerContributor.set(sub.contributorId, {
          contributor_id: sub.contributorId,
          html_content: sub.htmlContent,
          approved_at: sub.createdAt.toISOString(),
          percentage: sub.percentage || 0,
        })
      }
    }

    return NextResponse.json({
      success: true,
      data: Array.from(latestPerContributor.values()),
    })
  } catch (error: any) {
    console.error("GET /api/gallery error:", error)
    return NextResponse.json({ success: false, error: error.message || 'Internal Server Error' }, { status: 500 })
  }
}
