import { prisma } from '@/lib/prisma'
import { NextRequest, NextResponse } from 'next/server'

// GET /api/submissions?contributor_id=xxx
export async function GET(req: NextRequest) {
  const contributorId = req.nextUrl.searchParams.get('contributor_id')

  if (!contributorId) {
    return NextResponse.json({ error: 'contributor_id is required' }, { status: 400 })
  }

  const submissions = await prisma.submission.findMany({
    where: { contributorId },
    orderBy: { createdAt: 'desc' },
    select: {
      id: true,
      contributorId: true,
      analysisResultId: true,
      htmlContent: true,
      percentage: true,
      approved: true,
      reportJson: true,
      createdAt: true,
    },
  })

  return NextResponse.json(submissions)
}

// POST /api/submissions
export async function POST(req: NextRequest) {
  const body = await req.json()
  const { contributorId, analysisResultId, htmlContent, percentage, approved, reportJson } = body

  if (!contributorId || !htmlContent) {
    return NextResponse.json(
      { error: 'contributorId and htmlContent are required' },
      { status: 400 },
    )
  }

  const submission = await prisma.submission.create({
    data: { contributorId, analysisResultId, htmlContent, percentage, approved, reportJson },
  })

  return NextResponse.json(submission, { status: 201 })
}
