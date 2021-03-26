package swarm

func (meta *DownloadMeta) calculateBoundsForPiece(index int) (int, int) {
	begin := index * meta.PieceSize
	end := begin + meta.PieceSize
	if end > meta.FileSize {
		end = meta.FileSize
	}
	return begin, end
}

// just to handle case of last piece, else always returns pieceSize
func (meta *DownloadMeta) calculatePieceSize(index int) int {
	begin, end := meta.calculateBoundsForPiece(index)
	return end - begin
}
