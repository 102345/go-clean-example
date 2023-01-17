package productusecase

func (usecase usecase) Delete(id uint64) error {
	err := usecase.repository.Delete(id)

	if err != nil {
		return err
	}

	return nil
}
