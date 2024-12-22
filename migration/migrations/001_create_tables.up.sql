CREATE TABLE IF NOT EXISTS
    "word_audio" (
        "word_id" UUID NOT NULL,
        "path" VARCHAR NOT NULL,
        "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
        PRIMARY KEY ("word_id")
    );