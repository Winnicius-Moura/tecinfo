-- CreateTable
CREATE TABLE "Submission" (
    "id" TEXT NOT NULL PRIMARY KEY,
    "contributorId" TEXT NOT NULL,
    "analysisResultId" TEXT,
    "htmlContent" TEXT NOT NULL,
    "percentage" REAL,
    "approved" BOOLEAN,
    "createdAt" DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- CreateIndex
CREATE INDEX "Submission_contributorId_idx" ON "Submission"("contributorId");
