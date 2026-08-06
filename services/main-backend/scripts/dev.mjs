import { spawn, spawnSync } from "node:child_process";
import { join } from "node:path";

const rootDir = join(import.meta.dirname, "..");
const withGoEnv = join(rootDir, "scripts/with-go-env.mjs");

const startDependencies = () => {
  const result = spawnSync(
    "docker",
    ["compose", "up", "postgres", "redis", "-d", "--wait"],
    {
      cwd: rootDir,
      stdio: "inherit",
    }
  );

  if (result.status !== 0) {
    console.error("Failed to start postgres/redis. Is Docker Desktop running?");
    process.exit(result.status ?? 1);
  }
};

const main = () => {
  console.log("Starting postgres and redis...");
  startDependencies();

  const migration = spawnSync(
    process.execPath,
    [withGoEnv, "run", "./cmd/retail", "--migrate"],
    {
      cwd: rootDir,
      stdio: "inherit",
    }
  );

  if (migration.status !== 0) {
    process.exit(migration.status ?? 1);
  }

  const server = spawn(
    process.execPath,
    [withGoEnv, "run", "./cmd/retail"],
    {
      cwd: rootDir,
      stdio: "inherit",
    }
  );
  const worker = spawn(process.execPath, [withGoEnv, "run", "./cmd/worker"], {
    cwd: rootDir,
    stdio: "inherit",
  });

  const handleExit = function handleExit(code) {
    process.exit(code ?? 1);
  };

  server.on("exit", handleExit);
  worker.on("exit", handleExit);
};

main();
