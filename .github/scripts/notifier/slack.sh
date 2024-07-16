#!/usr/bin/env bash
set -eu

readonly TEST_RESULT_SUCCESS="success"
readonly TEST_RESULT_FAILURE="failure"
readonly TEST_RESULT_CANCELED="canceled"
readonly TEST_RESULT_SKIPPED="skipped"

color=""
message=""

read -r color message < <(
case "${TEST_RESULT}" in
  "${TEST_RESULT_SUCCESS}") echo "#74c7b8" "passed";;
  "${TEST_RESULT_FAILURE}") echo "#ef4f4f" "failed";;
  "${TEST_RESULT_CANCELED}") echo "#f4d160" "canceled";;
  "${TEST_RESULT_SKIPPED}") echo "#dddddd" "skipped";;
  *) echo "#dddddd" "unknown status: ${TEST_RESULT}"
esac
)

readonly COMMIT_HASH=$(echo "${GITHUB_SHA}" | cut -c 1-7)
readonly LINK=$(

if [[ "${GITHUB_REF}" =~ ^"refs/heads/" ]]; then
  readonly BRANCH=$(echo "${GITHUB_REF}" | cut -c 12-)
  echo "branch: <https://github.com/${GITHUB_REPOSITORY}/tree/${BRANCH}|${BRANCH}>"
else
  readonly TAG=$(echo "${GITHUB_REF}" | cut -c 11-)
  echo "tag: <https://github.com/${GITHUB_REPOSITORY}/releases/tag/${TAG}|${TAG}>"
fi
)

# Message Example:
# Workflow: CI (#7) of 42milez/go-oidc-expt (branch: the-first-step) passed.
# Commit: Apr. 10, 2023 (39e3068) by 42milez
readonly TEXT=$(
WF=${GITHUB_WORKFLOW}
REPO=${GITHUB_REPOSITORY}
RUN_ID=${GITHUB_RUN_ID}
RUN_NUM=${GITHUB_RUN_NUMBER}
ST_MSG=${message}
CMT_MSG=${COMMIT_MESSAGE}
SHA=${GITHUB_SHA}
CMT_HASH=${COMMIT_HASH}
ACT=${GITHUB_ACTOR}
cat <<EOF
Workflow: ${WF} (<https://github.com/${REPO}/actions/runs/${RUN_ID}|#${RUN_NUM}>) of <https://github.com/${REPO}|${REPO}> (${LINK}) ${ST_MSG}.
Commit: ${CMT_MSG} (<https://github.com/${REPO}/commit/${SHA}|${CMT_HASH}>) by <https://github.com/${REPO}/pulse|${ACT}>
EOF
)

readonly DATA=$(cat <<EOF
{
  "attachments": [
    {
      "blocks": [
        {
          "type": "section",
          "text": {
            "type": "mrkdwn",
            "text": "${TEXT}"
          }
        }
      ],
      "color": "${color}"
    },
  ],
  "blocks": [],
  "channel": "${SLACK_CHANNEL}",
  "text": "${TEST_RESULT^}",
  "username": "${SLACK_USERNAME}"
}
EOF
)

curl -s -X POST \
  -H "Content-type: application/json; charset=utf-8" \
  -H "Authorization: Bearer ${SLACK_BOT_USER_OAUTH_TOKEN}" \
  -d "${DATA}" \
  https://slack.com/api/chat.postMessage \
> /dev/null
