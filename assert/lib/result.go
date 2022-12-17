package lib

func AutoErrorType(err error) Status {
	if err == nil {
		return Ok
	}
	return Error
}

func AutoGuidance(err error) []string {
	if err != nil {
		return []string{err.Error()}
	}

	return nil
}

func ResultByError(check string, assertion string, err error, extraGuidance []string) Result {
	return Result{
		Status:    AutoErrorType(err),
		Check:     check,
		Assertion: assertion,
		Guidance:  append(AutoGuidance(err), extraGuidance...),
	}
}

func ToResult(status Status, check string, assertion string, guidance []string) Result {
	return Result{
		Status:    status,
		Check:     check,
		Assertion: assertion,
		Guidance:  guidance,
	}
}
