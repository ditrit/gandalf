#!/bin/bash

set -eu

MIN_PER_FUNC="75"
MIN_TOTAL="85"

COVER_MODE="atomic"

COVERAGE_FILE="coverage.out"
PROFILE_FILE="tmp_profile.out"
FINAL_OUT="tmp_coverage_final.out"
RES_FILE="tmp_res.out"

_RED=$(tput setaf 1 || echo "")
_GREEN=$(tput setaf 2 || echo "")
_YELLOW=$(tput setaf 3 || echo "")
# _BLUE=$(tput setaf 4 || echo "")
_PURPLE=$(tput setaf 5 || echo "")
# _CYAN=$(tput setaf 6 || echo "")
# _WHITE=$(tput setaf 7 || echo "")
_END=$(tput sgr0 || echo "")

OK_PROMPT="${_GREEN}OK${_END}"
KO_PROMPT="${_RED}==FAIL==${_END}"

echo "0" > "$RES_FILE"

(
	cd "$GOPATH/src/mia"
	echo "mode: ${COVER_MODE}" > "${COVERAGE_FILE}"

	COVERPKG=$(go list ./... | tr '\n' ',' | sed 's/,$//g')

	for d in $(go list ./... | grep -v vendor); do
		go test -covermode="${COVER_MODE}" -coverprofile="${PROFILE_FILE}" -coverpkg="${COVERPKG}" "$d"

		if [ -f ${PROFILE_FILE} ]; then
			tail -n +2 < ${PROFILE_FILE} >> "${COVERAGE_FILE}"
			rm -f ${PROFILE_FILE}
		fi
	done

	go tool cover -func="${COVERAGE_FILE}" > "${FINAL_OUT}"

	tr -s '\t' < "${FINAL_OUT}" | while read -r location nameFunc coverage;
	do
		minActu="$MIN_PER_FUNC"
		if [ "$location" = "total:" ]; then
			minActu="$MIN_TOTAL"
		fi
		PROMPT="$OK_PROMPT"
		if [ "${coverage%.*}" -lt "${minActu}" ]; then
			PROMPT="$KO_PROMPT"
			echo "1" > "${RES_FILE}"
		fi

		if [ "$location" = "total:" ]; then
			printf "\n"
			printf "${_PURPLE}%s${_END} %#+7s => %s\n" "$location" "$coverage" "$PROMPT"
			printf "\n"
		else
			printf "%-40s %-40s %#+7s => %s\n" "$location" "$nameFunc" "$coverage" "$PROMPT"
		fi
	done

	rm -f "${FINAL_OUT}"

	printf "\n"
	printf "Functions coverage must be greater than %s%s%%%s\n" "$_YELLOW" "$MIN_PER_FUNC" "$_END"
	printf "Total coverage must be greater than %s%s%%%s\n" "$_YELLOW" "$MIN_TOTAL" "$_END"
	printf "\n"
)

RES=$(cat "$RES_FILE")
rm -f $RES_FILE

exit "${RES}"
