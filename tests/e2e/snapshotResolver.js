// eslint-disable-next-line no-undef
let env = process.env.GO_DISCOVERY_E2E_ENVIRONMENT;
if (env != 'ci') {
  env = 'staging';
}

const snapshotDir = `tests/e2e/__snapshots__/${env}`;

// eslint-disable-next-line no-undef
module.exports = {
  // resolves from test to snapshot path
  resolveSnapshotPath: (testPath, snapshotExtension) =>
    testPath.replace('tests/e2e', snapshotDir) + snapshotExtension,

  // resolves from snapshot to test path
  resolveTestPath: (snapshotFilePath, snapshotExtension) =>
    snapshotFilePath.replace(snapshotDir, 'tests/e2e').slice(0, -snapshotExtension.length),

  // Example test path, used for preflight consistency check of the implementation above
  testPathForConsistencyCheck: 'tests/e2e/example.test.js',
};